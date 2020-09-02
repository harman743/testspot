# Andrew Vo

import sys
import os
print ("Welcome, \n")
fileName = input("Enter file name: ")

if not os.path.isfile(fileName):
    print("File not found.")
    sys.exit()


flagIn = input("Enter number of lines to skip at top and bottom (Press enter for default value of 5): ")

if flagIn == "":
    flag = 5

else:
    flag = int(flagIn)

with open(fileName) as fin:

    lines = fin.readlines()

    if len(lines) <= 2 * flag:
        for line in lines:
            print(line)
    else:
        for x in range(flag, len(lines) - flag):
            print(lines[x])


