package main;

import (
  "fmt";
  "math"
)

func main(){

  var isPrime bool;
  var sum int64 = 5;

  for i := 5; i < 20000; i += 2{

    isPrime = true;

    for j := 3; float64(j) < math.Floor(math.Sqrt(float64(i)))+1; j += 2 {
      if(i % j == 0){
        isPrime = false;
        break;
      }
    }

    if(isPrime == true){
      sum += int64(i);
    }


  }

  fmt.Printf("%d\n", sum);

}
