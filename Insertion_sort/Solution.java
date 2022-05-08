class Solution {
    public static int[] insert(int[] array, int rightIndex, int value) {
      // write this method
      int endpos = 0;
      for(int i=rightIndex;i<array.length;i++){
        if (array[i]==value){
          endpos = i;
          break;
        }
      }
  
      for (int i=0;i<=rightIndex;i++){
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
      return array;
    } 

    public static void main(String[] args) {
        String res = "";  
        int[] result = Solution.insert(new int[]{3, 5, 7, 11, 13, 2, 9, 6},4,2);
          for (int i=0 ; i< result.length; i++){
              res += result[i] + ",";
          }
          System.out.println(res);
    }
}