from django.db import models

class Artist(models.Model):
    #アーティスト名
    name = models.CharField(max_length=50)
    #代表作品名
    discography = models.CharField(max_length=200)
    #活動開始時期
    debut_date = models.DateField()
    #レーベル
    agency = models.CharField(max_length=50)
    #音楽ジャンル
    GENRE_CHOICES = [
        ('pop', 'Pop'),
        ('rock', 'Rock'),
        ('metal', 'Metal'),
        ('hard rock', 'Hard Rock'),
        ('progressive', 'Progressive'),
    ]
    genre = models.CharField(max_length=50, choices=GENRE_CHOICES, null=True, blank=True)