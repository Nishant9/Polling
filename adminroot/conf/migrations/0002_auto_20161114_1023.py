# -*- coding: utf-8 -*-
# Generated by Django 1.10 on 2016-11-14 10:23
from __future__ import unicode_literals

from django.db import migrations


class Migration(migrations.Migration):

    dependencies = [
        ('conf', '0001_initial'),
    ]

    operations = [
        migrations.RenameModel(
            old_name='BulkVoters',
            new_name='BulkVoter',
        ),
        migrations.RenameModel(
            old_name='Voters',
            new_name='Voter',
        ),
    ]
