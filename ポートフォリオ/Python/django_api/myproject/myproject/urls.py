from django.urls import path
from myapp.views import ArtistListCreate

urlpatterns = [
    path('artists/', ArtistListCreate.as_view()),
]