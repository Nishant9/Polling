from django.contrib import admin

# Register your models here.
from .models import *

admin.site.register(Voter)
admin.site.register(Candidate)
admin.site.register(BulkVoter)
