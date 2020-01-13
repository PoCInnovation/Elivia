#!/usr/bin/env python
# -*- coding:utf-8 -*-

import utils
import webbrowser

def run(string, entities):
	"""Launch web-browser"""

	domains = []
	#output = ''
	i = 0

	for item in entities:
		if item['entity'] == 'url':
			domains.append(item['resolution']['value'].lower())
			while i < len(domains):
				webbrowser.open(domains[i], new=2)
				i += 1
		else:
			return utils.output('end', 'invalid_domain_name', utils.translate('invalid_domain_name'))
	return utils.output('end', 'done', utils.translate('done'))