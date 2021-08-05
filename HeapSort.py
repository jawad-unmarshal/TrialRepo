def heapify(arr, n , root):

    smol = root

    leftNode = 2*root + 1
    RightNode = 2*root + 2

    if leftNode < n and arr[root] > arr[leftNode]:
        smol = leftNode
        arr[root],arr[smol] = arr[smol],arr[root]
       

    if RightNode < n and arr[smol] > arr[RightNode]:
        smol = RightNode
        arr[root],arr[smol] = arr[smol],arr[root]
        
    if smol != root:        
        heapify(arr,n,smol)
    
    

def heapsort(arr):

    n = len(arr)
    for i in range(n//2-1,-1,-1):
        heapify(arr,n,i)
    
    for i in range(n-1, 0, -1):
    
        arr[i],arr[0] = arr[0],arr[i]
        heapify(arr,i,0)

from os import system
system('clear')

print("Enter space seperated values for an array:")
MyArr = list(map(int,input().split(" ")))
heapsort(MyArr)
print(MyArr)
