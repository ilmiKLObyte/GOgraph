package main
import "fmt"

func classic(){
    var userInputInt uint16
    var userInputTemp uint8
    var i uint16
    var j uint16
    var total uint16
    //var userInputTemp2 uint8
    fmt.Println("How many items:")
    fmt.Scanf("%d",&userInputInt)
    mainArray := make([]uint16,0)
    
    for i = 0;i<userInputInt;i++{
        mainArray = append(mainArray,0)
    }
    stringArray := make([]string,userInputInt)
    fmt.Println("Give each item a name:")
    for i=0;i<userInputInt;i++{
        fmt.Scanf("%s",&stringArray[i])
    }
    for(1<10){
        fmt.Println("Chose an option: ")
        fmt.Println("1.Add")
        fmt.Println("2.Remove")
        fmt.Println("3.See graph and data")
        fmt.Scanf("%d",&userInputTemp)
        switch userInputTemp {
            case 1:
                fmt.Println("Which position? ",userInputInt)
                fmt.Scanf("%d",&userInputTemp)
                userInputTemp--
                total++
                mainArray[userInputTemp] = mainArray[userInputTemp] + 1
            case 2:
                fmt.Println("Which position? ",userInputInt)
                fmt.Scanf("%d",&userInputTemp)
                userInputTemp--
                mainArray[userInputTemp] = mainArray[userInputTemp] - 1
            case 3:
                for i = 0; i < userInputInt; i++{
                    fmt.Print(i+1,"|")
                    for j = 0; j < mainArray[i]; j++{
                        fmt.Print(" # ")
                        
                    }
                    fmt.Print(stringArray[i])
                    fmt.Println()
                }
                //fmt.Println(total)
                fmt.Println(mainArray)
                fmt.Println("percentage:")
                for i=0;i<userInputInt;i++{
                    var porCem float32 = float32(mainArray[i]) / float32(total)
                    fmt.Println(porCem*100,"%")
                }
        }
    }
}


func main(){
    var userInputInt uint8
    fmt.Println("Modes:\n1.Classic")
    fmt.Scanf("%d",&userInputInt)
    switch userInputInt {
        case 1:
            classic()
        default:
            fmt.Println("Not known.")
            main()
    }
}


