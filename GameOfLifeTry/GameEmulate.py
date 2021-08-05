def WillAlive(Location, PosOccupied,Min=2):
    '''This will let the system know if a node will be alive or not for the next iteration'''
    # To stay alive Needs 2 or 3 neighbors. More or less dies
    # If Min is 3 then it will rise from the dead.
    NeighBors = [
        (Location[0]-1,Location[1]-1),(Location[0]-1,Location[1]),(Location[0]-1,Location[1]+1),
        (Location[0],Location[1]-1),(Location[0],Location[1]+1),
        (Location[0]+1,Location[1]-1),(Location[0]+1,Location[1]),(Location[0]+1,Location[1]+1)
        ]
    NumNeigh = 0
    for Neighbour in NeighBors:
        if (Neighbour[0] > 0 and Neighbour[1]>0):
            if Neighbour in PosOccupied:
                NumNeigh += 1
    if NumNeigh >= Min and NumNeigh <= 3:
        return True
    else:
        return False
        


def NeighbourAlive(Location,CurrentPosition,CurrentStack):
    '''
    This function will check if a dead neighbour will come alive.
    Ever node can have a maximum of four dead neighbours it's directly in contact with it possibly  come alive
    '''
    
    AliveNeighbors = [
        (Location[0]-1,Location[1]),
        (Location[0],Location[1]-1),(Location[0],Location[1]+1),
        (Location[0]+1,Location[1])
    ]
    for Nbr in AliveNeighbors:
        if(Nbr[0] > 0 and Nbr[1]>0 and Nbr not in CurrentPosition):
            if(WillAlive(Nbr,CurrentPosition,3)):
                CurrentStack.add(Nbr)
    return CurrentStack




# These are the start positions
# StartPos = {(2,1),(2,2),(2,3)}
# CurPos = StartPos.copy()

def SetUpStartState():
    '''This function is for taking input initially'''
    print("Enter ',' seperated co-ordinates: (Enter Q when done)")
    Coords = set()
    while True:
        Val = input()        
        if(Val != 'Q' and Val != 'q'):
            item = tuple(int(x) for x in Val.split(','))
            # print(item)
            Coords.add(item)
        else:
            break
    # print(Coords)
    return Coords


def ComputeItr(Positions):
    '''We compute the result of an iteration here. The function returns the next iteration once done'''
    NewStat = set()
    for Loc in Positions:
        if WillAlive(Loc,Positions,2):
            NewStat.add(Loc)
        NewStat = NeighbourAlive(Loc,Positions,NewStat)
    return NewStat

def main():
    StartPos = SetUpStartState()
    runNo = 0
    while True:
        runNo += 1
        print("Result after run {0}".format(runNo))
        StartPos = ComputeItr(StartPos)
        if StartPos:
            print(StartPos)
        else:
            print("Empty Grid")
        print("Enter 'y' or 'Y' to continue anything else to exit")
        ch = input()
        if(not(ch == 'y' or ch == 'Y')):
            break
from os import system
if __name__ == '__main__':
    system('clear')
    main()
# main()


