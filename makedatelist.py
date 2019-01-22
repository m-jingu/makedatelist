#!/usr/bin/env python

import sys
import argparse
from datetime import datetime, timedelta

version = '%(prog)s 1.0'

def ArgParse():
    parser = argparse.ArgumentParser(description=
    '''This script make a list of range of specific two dates.
    
    Create Date: 2016-07-25 ''',
    formatter_class=argparse.RawDescriptionHelpFormatter)
    
    parser.add_argument('-s', '--startdate', action='store', metavar='DATE', dest='sdate', type=str, help='start date - format YYYYMMDD - default 20160101', default='20160101')
    parser.add_argument('-e', '--enddate', action='store', metavar='DATE', dest='edate', type=str, help='end date - format YYYYMMDD - default 20160131', default='20160131')
    parser.add_argument('-f', '--format', action='store', type=str, help='format - default %%Y-%%m-%%d', default='%Y-%m-%d')
    parser.add_argument('-v', '--version', action='version', version=version)
    args = parser.parse_args()

    return args

def daterange(start_date, end_date):
    for n in range((end_date - start_date).days):
        yield start_date + timedelta(n)

if __name__ == "__main__":
    args = ArgParse()
    try:
        start = datetime.strptime(args.sdate, '%Y%m%d')
        end = datetime.strptime(args.edate, '%Y%m%d')
    except ValueError:
        print("incorrect timestamp")
        sys.exit()
    except:
        raise

    for i in daterange(start, end):
        print(i.strftime(args.format))

