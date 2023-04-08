from django.db import models

class Artist(models.Model):
    name = models.CharField(max_length=200)
    debut_date = models.DateField()
    agency = models.CharField(max_length=200)