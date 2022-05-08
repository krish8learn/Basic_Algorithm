

public class SS {
    static void swap(int[] array, int firstIndex, int secondIndex) {
        int temp = array[firstIndex];
        array[firstIndex] = array[secondIndex];
        array[secondIndex] = temp;
        return;
      }
      
      static int indexOfMinimum(int[] array, int startIndex) {
        int minValue = array[startIndex];
        int minIndex = startIndex;
    
        for(int i = minIndex + 1; i < array.length; i++) {
          if(array[i] < minValue) {
            minIndex = i;
            minValue = array[i];
          }
        } 
        return minIndex;
      }
    
      public static int[] selectionSort(int[] array) {
        // Write this method
        for(int i=0; i<array.length-1; i++){
          int secondIndex = indexOfMinimum(array, i);
          swap(array,i,secondIndex);
        }
        return array; 
      }
  
      public static void main(String[] args) {
        String res = "";  
        int[] result = SS.selectionSort(new int[]{4, 3, 2, 1, 0, -1, -99});
          for (int i=0 ; i< result.length; i++){
              res += result[i] + ",";
          }
          System.out.println(res);
      }
}
