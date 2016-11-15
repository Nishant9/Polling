from django.db import models

# Create your models here.
from django.db import models
from django.conf import settings

class Candidate(models.Model):
    name = models.CharField(max_length = 120)
    email = models.CharField(max_length = 120, null = True, blank = True)
    photo = models.CharField(max_length = 120, null = True, blank = True)

    def __str__(self):
        return str(self.name) + " " + str(self.email)

class Voter(models.Model):
    name = models.CharField(max_length = 120)
    email = models.CharField(max_length = 120, null = True, blank = True)
    passwd = models.CharField(max_length = 120, null = True, blank = True)
    def __str__(self):
        return str(self.name) + " " + str(self.email)


class BulkVoter(models.Model):
    tag = models.CharField(max_length = 50)
    docfile = models.FileField(upload_to = 'documents/')
    def __str__(self) :
        return str(self.tag) + "  " + str(self.docfile)
