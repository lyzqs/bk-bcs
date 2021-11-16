# -*- coding: utf-8 -*-
"""
Tencent is pleased to support the open source community by making 蓝鲸智云PaaS平台社区版 (BlueKing PaaS Community
Edition) available.
Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://opensource.org/licenses/MIT

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.
"""
# Generated by Django 1.11.23 on 2020-07-02 11:34
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('configuration', '0034_auto_20200106_1455'),
    ]

    operations = [
        migrations.AlterField(
            model_name='resourcefile',
            name='name',
            field=models.TextField(verbose_name='file name'),
        ),
        migrations.AlterField(
            model_name='resourcefile',
            name='resource_name',
            field=models.CharField(choices=[('Deployment', 'Deployment'), ('Service', 'Service'), ('ConfigMap', 'ConfigMap'), ('Secret', 'Secret'), ('Ingress', 'Ingress'), ('StatefulSet', 'StatefulSet'), ('DaemonSet', 'DaemonSet'), ('Job', 'Job'), ('HPA', 'HPA'), ('ServiceAccount', 'ServiceAccount'), ('ClusterRole', 'ClusterRole'), ('ClusterRoleBinding', 'ClusterRoleBinding'), ('PodDisruptionBudget', 'PodDisruptionBudget'), ('StorageClass', 'StorageClass'), ('PersistentVolume', 'PersistentVolume'), ('PersistentVolumeClaim', 'PersistentVolumeClaim'), ('CustomManifest', 'CustomManifest')], max_length=32),
        ),
    ]