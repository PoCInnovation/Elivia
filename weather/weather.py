#!/usr/bin/env python
# -*- coding:utf-8 -*-

import utils
import sys
import difflib
import requests
import packages.weather.cities as cities

def weather(string, entities):
    testcity = ["not found"]
    elems = string.split()
    for e in elems:
        testcity.append(difflib.get_close_matches(e, cities.cities, n=1, cutoff=0.8))
    city = ""
    for i in testcity:
        if len(i) > 0:
            if len(i[0]) > 3:
                city = i[0]
    if city == "":
        return utils.output('end', 'notfound', utils.translate('notfound'))
    lt = cities.cities[city][0]
    lg = cities.cities[city][1]
    req = requests.get(f"https://www.metaweather.com/api/location/search/?lattlong={lt},{lg}")
    woeid = req.json()[0]["woeid"]
    req = requests.get(f"https://www.metaweather.com/api/location/{woeid}/")
    resp = req.json()
    temp = resp["consolidated_weather"][0]["the_temp"]
    return utils.output('end', 'found', utils.translate('found', {'temperature': temp, 'city':city}))
