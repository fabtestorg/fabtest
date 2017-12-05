#!/usr/bin/python
#coding=utf-8 
import os

def loaddata(file):
	try:
		fo = open(file, "r")
	except IOError:
		print "文件不存在！"
		exit()
	results = []
	for line in fo.readlines():
		try:
			n = int(line)
		except Exception:
			print "格式错误"
		else:
			results.append(n)
	return results

def statistics(filename, results):
	less1 = 0 # 小于1s的
	less3 = 0 # 小于3s的
	more3 = 0 # 大于3s的
	more8 = 0 # 大于8s的
	total = 0
	length = len(results)
	for i in results:
		total += i
		if i < 1000:
			less1+=1
			less3+=1
		elif i <= 3000:
			less3+=1
		else:
			more3+=1
			if i >= 8000:
				more8+=1

	print ("%s: aver:%.2f,max:%d,min:%d,total:%d    <3s:%d,%.2f    >3s:%d,%.2f    <1s:%d,%.2f    >8s:%d,%.2f"%(filename,float(total)/length,max(results),min(results),length,less3,float(less3)/length,more3,float(more3)/length,less1,float(less1)/length,more8,float(more8)/length))


def scan_files(directory, prefix=None, postfix=None):
	files_list = []
	for root, sub_dirs, files in os.walk(directory):
		for special_file in files:
			if postfix:
				if special_file.endswith(postfix):
					files_list.append(os.path.join(root, special_file))
			elif prefix:
				if special_file.startswith(prefix):
					files_list.append(os.path.join(root, special_file))
			else:
				files_list.append(os.path.join(root, special_file))
	return files_list

def analyse(config_dir):
	directory = "%sevent_logs"%config_dir
	postfix = ".log"
	print "扫描目录为:", directory
	print "后缀名为:", postfix
	filelist = scan_files(directory, None, postfix)
	for file in filelist:
		results = loaddata(file)
		statistics(file, results)

