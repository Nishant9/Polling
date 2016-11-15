from django.shortcuts import render
from django.http import HttpResponse

from django.shortcuts import redirect, render
from django.template import loader, Context, Template
from django.http import HttpResponseRedirect
from .models import BulkVoter, Voter, Candidate
from django.template import TemplateDoesNotExist
from django.http import Http404
from django.views.static import serve
import os


# Create your views here.


def index(request):
    template = loader.get_template('index.html')
    return HttpResponse(template.render())


def config(request):
    config_file = '''
package election_conf

var List = map[string][]string{
'''
    voter_list = [t[0] for t in list(Voter.objects.all().values_list('name'))]
    if len(voter_list) > 0 :
        config_file += '"manual": []string{ '
        config_file += str(voter_list)[1:-1].replace('\'','"')
        config_file += ',},\n'

    bulkvoter_list = list(BulkVoter.objects.all().values_list('tag', 'docfile', 'id'))

    if len(bulkvoter_list) > 0 :
        import csv
        for tag, filename, i in bulkvoter_list :
            voter_list = []
            with open(filename, 'r') as f :
                reader = csv.reader(f)
                row = next(reader)
                if row[0] != 'name' and row[0] != 'username' :
                    voter_list += row[0]
                    if len(row) == 1 :
                        temp_voter = Voter(name = row[0])
                    elif len(row) == 2 :
                        temp_voter = Voter(name = row[0], email = row[1])
                    elif len(row) == 3 :
                        temp_voter = Voter(name = row[0], email = row[1], passwd = row[2])
                    temp_voter.save()
                for row in reader :
                    voter_list.append(row[0])
                    if len(row) == 1 :
                        temp_voter = Voter(name = row[0])
                    elif len(row) == 2 :
                        temp_voter = Voter(name = row[0], email = row[1])
                    elif len(row) == 3 :
                        temp_voter = Voter(name = row[0], email = row[1], passwd = row[2])
                    temp_voter.save()
            if len(voter_list) > 0 :
                config_file +=  ('"' +tag + '" : []string{ ')
                config_file += str(voter_list)[1:-1].replace('\'','"')
                config_file += '},\n'
            BulkVoter.objects.get(id = i).delete()

    config_file += '}\n\n'


    candidate_list = [t for t in list(Candidate.objects.all().values_list('name','photo'))]
    if len(candidate_list) > 0 :
        config_file += 'var Candidates = map[string]string{\n'
        for cd,ph in candidate_list :
            config_file += ( '"' + cd + '" : ' + '"' + ph + '",\n' )
        config_file += '}\n'


    config_file += '''
const Number_of_votes int = 1
const Pass_Length int =8
    '''

    response = HttpResponse(config_file, content_type='text/plain')
    response['Content-Disposition'] = 'attachment; filename=election_conf.go'
    return response

import string
import random

def pass_generator(size=8, chars=string.ascii_uppercase + string.digits + string.ascii_lowercase):
    return ''.join(random.choice(chars) for _ in range(size))

def mail_passwords(list) :
    import sys
    import os
    import re
    import time
    import sched
    import random
    #from smtplib import SMTP_SSL as SMTP
    from smtplib import SMTP
    from email.mime.text import MIMEText

    scheduler = sched.scheduler(time.time, time.sleep)


    def tasker(info) :
        SMTPserver = 'smtp.cc.iitk.ac.in'
        sender =     'nishgu@iitk.ac.in'
        USERNAME = "nishgu"
        PASSWORD = "Tiaspei2"

        # typical values for text_subtype are plain, html, xml
        text_subtype = 'plain'

        content="""\
Hello,
The password for user {username} is {password}.
Thank You
ElectionBot
"""

        subject="Password for Election"

        msg = MIMEText(content.format(username = info[0], password = info[2]), text_subtype)
        msg['Subject']= subject
        msg['From']   = sender # some SMTP servers will do this automatically, not all
        try:
            conn = SMTP(SMTPserver)
            conn.set_debuglevel(False)
            conn.login(USERNAME, PASSWORD)

            try:
                conn.sendmail(sender, [info[1]], msg.as_string())
            except:
                print("Could not send mail to",destination)
        except Exception as exc:
            print( "mail failed; %s" % str(exc) ) # give a error message
        finally :
            conn.quit()


    for info in list:
        randintg = random.randint(1,100)
        print(randintg,' ',info[0], info[1])
        scheduler.enter(randintg,1,tasker,(info,))
    scheduler.run()


import threading
def mail(request):
    voter_list = list(Voter.objects.all().values_list('name','email','passwd'))
    #mail_passwords(voter_list)
    mail_thread = threading.Thread(target=mail_passwords, args=(voter_list,))
    mail_thread.start()
    return redirect(index)


def passwd(request):
    for voter  in Voter.objects.all() :
        if voter.passwd == '' or voter.passwd == None :
            voter.passwd = pass_generator()
            voter.save()
    sql = '''
drop database poll_1;
create database poll_1;
use poll_1;
create table authdb (username VARCHAR(20) UNIQUE KEY PRIMARY KEY, passwords TEXT);
create table ballot (username VARCHAR(50) PRIMARY KEY, vote_0 VARCHAR(40)); -- add as many as vote_i
'''
    insert_holder = 'INSERT into authdb (username,passwords) VALUES (\'{user}\', \'{passwd}\');\n'
    voter_list = list(Voter.objects.all().values_list('name','passwd'))
    for us, pwd in voter_list :
        sql += insert_holder.format(user = us,passwd = pwd)
    response = HttpResponse(sql, content_type='text/plain')
    response['Content-Disposition'] = 'attachment; filename=sql_setup.sql'
    return response

