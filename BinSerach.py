def BinS(arr,key):
    UpLim, LLim = len(arr), 0
    while(UpLim > LLim):
        mid = (UpLim+LLim)//2
        if arr[mid] == key:
            return True
        if key < arr[mid]:
            UpLim = mid-1
        else:
            LLim = mid+1
    return False

class MySorter:
    def __init__(self,ar):
        self.ar = MySorter.Splitter(ar)

    def Splitter(MyAr):
        if(len(MyAr)>1):
            Mid = len(MyAr)//2
            L = MyAr[:Mid]
            R = MyAr[Mid:]
            MySorter.Splitter(L)
            MySorter.Splitter(R)
            MySorter.Combiner(MyAr,L,R)        

    def Combiner(MyAr,L,R):
        i,j,k = 0,0,0
        while i < len(L) and j < len(R):

            if L[i]<R[j]:
                MyAr[k] = L[i]
                i += 1
                
            else:
                MyAr[k] = R[j]
                j += 1
            k += 1

        while(i < len(L)):
            MyAr[k] = L[i]
            k += 1
            i += 1
        while(j < len(R)):
            MyAr[k] = R[j]
            k += 1
            j += 1
    # return self.arr
    # print (arr)
from os import system as sys
sys("clear")
print("Enter space seperated values for an array:")
MyArr = list(map(int,input().split(" ")))
key = int(input("Enter the key to find in the array: "))
MySorter(MyArr)
if BinS(MyArr,key):
    print("{0} was found in the array".format(key))
else:
    print("{0} was not found in the array".format(key))




