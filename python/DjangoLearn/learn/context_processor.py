# coding:utf-8

from django.conf import settings

def setting(request):
    return {'setting':settings}

def ip_address(request):
    return {'ip_address':request.META['REMOTE_ADDR']}