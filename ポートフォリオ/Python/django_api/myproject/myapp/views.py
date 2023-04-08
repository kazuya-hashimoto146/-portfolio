from rest_framework.response import Response
from rest_framework import generics
from .models import Artist
from .serializers import ArtistSerializer

class ArtistListCreate(generics.ListCreateAPIView):
    queryset = Artist.objects.all()
    serializer_class = ArtistSerializer
    template_name = 'artists.html'

    def list(self, request, *args, **kwargs):
        queryset = self.filter_queryset(self.get_queryset())
        serializer = self.get_serializer(queryset, many=True)
        return Response({'artists': serializer.data})