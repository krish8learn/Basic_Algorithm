package JAVA;
public class IS {
    static void insert(int[] array, int rightIndex, int value) {
        int endpos = 0;
          for(int i=rightIndex;i<array.length;i++){
            if (array[i]==value){
              endpos = i;
              break;
            }
          }
      
        for(int i=0;i<=rightIndex;i++){
            if(value<array[i]){
              for(int j=endpos;j>i;j--){
                int k =j; 
                array[j] = array[--k];
                //System.out.println(array[j]+","+j+","+k+","+array[k]);
              }
              array[i]= value;
              break;
            }
        }
        return;
    }

    public static int[] insertionSort(int[] array) {
        //Write this method
        for(int i=0;i<array.length-1;i++){
          insert(array,i+1,array[i+1]);
        } 
        return array;
    }
    public static void main(String[] args) {
        int[] array = {4, 3, 2, 1, 0, -1, -99};
        int [] result =IS.insertionSort(array);
        String res = "";
        System.out.println("END");
        for(int i=0; i<result.length; i++){
            res += result[i] + ",";
        }
        System.out.println(res);
    }
}
