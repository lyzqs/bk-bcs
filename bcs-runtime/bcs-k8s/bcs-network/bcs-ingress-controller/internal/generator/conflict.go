/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.,
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package generator

import (
	"context"
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	networkextensionv1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/apis/networkextension/v1"
)

func generateRuleConflictMessage(rule networkextensionv1.IngressRule,
	ingressName, ingressNamespace, lisIngressName, lisIngressNamespace string) string {
	return fmt.Sprintf("[conflict] rule %+v of ingress %s/%s is conflict with ingress %s/%s",
		rule, ingressName, ingressNamespace, lisIngressName, lisIngressNamespace)
}

func generateMappingConflictMessage(mapping networkextensionv1.IngressPortMapping,
	ingressName, ingressNamespace, lisIngressName, lisIngressNamespace string) string {
	return fmt.Sprintf("[conflict] mapping %+v of ingress %s/%s is conflict with ingress %s/%s",
		mapping, ingressName, ingressNamespace, lisIngressName, lisIngressNamespace)
}

func getKeyByValue(m map[string]string, value string) string {
	for k, v := range m {
		if v == value {
			return k
		}
	}
	return ""
}

func isRuleConflict(lbID, ingressName, ingressNamespace string,
	rule networkextensionv1.IngressRule,
	existedListenerMap map[string]networkextensionv1.Listener) (bool, string) {

	tmpKey := GetListenerName(lbID, rule.Port)
	existedListener, ok := existedListenerMap[tmpKey]
	if !ok {
		return false, ""
	}

	// when existed listener with same port was generated by this ingress, we think it is not conflict
	// otherwise, return true and return conflict message
	ingressNameValue, okLisIngressName := existedListener.Labels[ingressName]
	if !okLisIngressName || ingressNameValue != networkextensionv1.LabelValueForIngressName ||
		ingressNamespace != existedListener.GetNamespace() {
		conflictName := getKeyByValue(existedListener.Labels, networkextensionv1.LabelValueForIngressName)
		conflictNs := existedListener.GetNamespace()
		conflictMsg := generateRuleConflictMessage(rule, ingressName, ingressNamespace,
			conflictName, conflictNs)
		blog.Warnf(conflictMsg)
		return true, conflictMsg
	}
	return false, ""
}

func isMappingConflict(lbID, ingressName, ingressNamespace string,
	mapping networkextensionv1.IngressPortMapping,
	existedListenerMap map[string]networkextensionv1.Listener) (bool, string) {

	segmentLen := mapping.SegmentLength
	if segmentLen == 0 {
		segmentLen = 1
	}
	istart := mapping.StartPort + mapping.StartIndex*segmentLen
	iend := mapping.StartPort + mapping.EndIndex*segmentLen
	for i := istart; i < iend; i++ {
		tmpKey := GetListenerName(lbID, i)
		existedListener, ok := existedListenerMap[tmpKey]
		if !ok {
			continue
		}
		ingressNameValue, okLisIngressName := existedListener.Labels[ingressName]
		if !okLisIngressName || ingressNameValue != networkextensionv1.LabelValueForIngressName ||
			ingressNamespace != existedListener.GetNamespace() {
			conflictName := getKeyByValue(existedListener.Labels, networkextensionv1.LabelValueForIngressName)
			conflictNs := existedListener.GetNamespace()
			conflictMsg := generateMappingConflictMessage(mapping, ingressName, ingressNamespace,
				conflictName, conflictNs)
			blog.Warnf(conflictMsg)
			return true, conflictMsg
		}
	}
	return false, ""
}

// return true, if the ingress is conflicts with existed listener
func (g *IngressConverter) checkConflicts(lbID string, ingress *networkextensionv1.Ingress) (bool, error) {
	existedListeners := &networkextensionv1.ListenerList{}
	err := g.cli.List(context.TODO(), existedListeners, &client.ListOptions{})
	if err != nil {
		blog.Errorf("failed list existed Listeners err %s", err.Error())
		return false, fmt.Errorf("failed list existed Listeners err %s", err.Error())
	}

	// use lbid-port as key of map for check conflicts
	existedListenerMap := make(map[string]networkextensionv1.Listener)
	for index, listener := range existedListeners.Items {
		// listener for port segment
		if listener.Spec.EndPort > 0 {
			for i := listener.Spec.Port; i <= listener.Spec.EndPort; i++ {
				tmpKey := GetListenerName(listener.Spec.LoadbalancerID, i)
				existedListenerMap[tmpKey] = existedListeners.Items[index]
			}
			continue
		}
		tmpKey := GetListenerName(listener.Spec.LoadbalancerID, listener.Spec.Port)
		existedListenerMap[tmpKey] = existedListeners.Items[index]
	}

	for _, rule := range ingress.Spec.Rules {
		isConflict, _ := isRuleConflict(lbID, ingress.GetName(), ingress.GetNamespace(), rule, existedListenerMap)
		if isConflict {
			return true, nil
		}
	}
	for _, mapping := range ingress.Spec.PortMappings {
		isConflict, _ := isMappingConflict(lbID, ingress.GetName(), ingress.GetNamespace(), mapping, existedListenerMap)
		if isConflict {
			return true, nil
		}
	}
	return false, nil
}
