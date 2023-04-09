from django.contrib import admin
from django.urls import path
from myapp.views import ArtistListCreate

urlpatterns = [
    path('admin/', admin.site.urls),
    path('artists/', ArtistListCreate.as_view()),
]