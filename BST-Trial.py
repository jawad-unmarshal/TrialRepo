from os import system
class BinTree:
    NumNodes = 0

    def __init__(self,val):
        self.left = None
        self.right = None
        self.data = val
        BinTree.NumNodes += 1
        print("{0} was added to the tree.\nIt Currently has {1} nodes".format(val,BinTree.NumNodes))


    def InsNode(self,val):
        if val == self.data:
            print("{0} Already exists in the tree".format(val))
            return
        if(val < self.data):
            # Val is lesser than original
            if(self.left == None):
                self.left = BinTree(val)
            else:
                self.left.InsNode(val)
             # Val is lesser than original
        else:
            if(self.right == None):
                self.right = BinTree(val)
            else:
                self.right.InsNode(val)
    def FindNode(self,val):
        if val ==  self.data:
            print("FOUND")
            return
        elif(val < self.data):
            if(self.left != None):
                self.left.FindNode(val)
            else:
                print("{0} does not exist in the tree".format(val))
                return
        else:
            if(self.right != None):
                self.right.FindNode(val)
            else:
                print("{0} does not exist in the tree".format(val))
                return
    def DelMyNode(RootNode, value):
        if(RootNode == None):
            print("Did not Find Node")
            return None
        if(value < RootNode.data):
           RootNode.left =  BinTree.DelMyNode(RootNode.left,value)           
        elif(value > RootNode.data):
            RootNode.right = BinTree.DelMyNode(RootNode.right,value)
        else:
            if(RootNode.left == None and RootNode.right ==  None):
                return None
            elif(RootNode.left == None):
                return RootNode.right
            elif(RootNode.right == None):
                return RootNode.left
            else:
                # Both have value, so move what's on the right up and make it the root.
                RemNode = RootNode.right
                MinVal = RemNode.data
                while RemNode.left:
                    RemNode = RemNode.left
                    MinVal = RemNode.data
                RootNode.data = MinVal
                RootNode.right = BinTree.DelMyNode(RootNode.right,MinVal)
                # return RemNode
        return RootNode
    def Preorder(node,prStat=False): 
        if not node: 
            return 
        if(prStat):     
            print(node.data)
        else:
            BinTree.NumNodes += 1
        BinTree.Preorder(node.left,prStat) 
        BinTree.Preorder(node.right,prStat) 

    def DelNode(Root,Val):
        BinTree.NumNodes = 0
        Root = BinTree.DelMyNode(Root,Val)
        BinTree.Preorder(Root)
        return Root
    

def main():
    
    while True:        
        print("\n1. Add Node\n2. Find Node\n3. Delete Node\n4. Preorder\n5.Exit\n\nEnter choice:-")
        choice = int(input())
        if(choice <= 3):

            val = int(input("Enter the value: "))
            if choice == 1:
                if(BinTree.NumNodes == 0):
                    Root = BinTree(val)
                else:
                    Root.InsNode(val)
            elif choice == 2:
                if(BinTree.NumNodes == 0):
                    print("Nothing in tree")
                else:
                    Root.FindNode(val)
            elif choice == 3:
                if(BinTree.NumNodes == 0):
                    print("Nothing in tree")
                else:
                    Root = BinTree.DelNode(Root,val)
                    print("{0} Node/s left in the tree".format(BinTree.NumNodes))
        elif choice == 4:
            print("\nPreOrder: ")
            BinTree.Preorder(Root, True)

        else:

            print("Byyeeeee")
            break

if __name__ == "__main__":
    system('clear')
    main()
# main()

