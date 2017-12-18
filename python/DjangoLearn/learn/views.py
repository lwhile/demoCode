# coding:utf-8

from django.shortcuts import render
from django.http import HttpResponse
# Create your views here.

# 创建我们的表单类
from .froms import AddForm
import json


def index(request):
    return render(request,'index.html')

def columns(request):
    return render(request,'index.html')


def add(request):
    a = request.GET['a']
    b = request.GET['b']
    a = int(a)
    b = int(b)
    return HttpResponse(str(a + b))
