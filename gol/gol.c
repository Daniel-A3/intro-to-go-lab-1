/*any live cell with fewer than two live neighbours dies
any live cell with two or three live neighbours is unaffected
any live cell with more than three live neighbours dies
any dead cell with exactly three live neighbours becomesalive*/

void updateBoard(uchar ** A, uchar ** B, int IMWD, int IMHT)
{
  int y,x,sum;

   for( y = 0; y < IMHT; y++ ) {
    for( x = 0; x < IMWD; x++ ) {
      sum = A[(y+IMHT-1)%IMHT][(x+IMWD-1)%IMWD] + A[(y+IMHT-1)%IMHT][( x+IMWD )%IMWD] + A[(y+IMHT-1)%IMHT][(x+IMWD+1)%IMWD] + \
            A[( y+IMHT )%IMHT][(x+IMWD-1)%IMWD]                    +                    A[( y+IMHT )%IMHT][(x+IMWD+1)%IMWD] + \
            A[(y+IMHT+1)%IMHT][(x+IMWD-1)%IMWD] + A[(y+IMHT+1)%IMHT][( x+IMWD )%IMWD] + A[(y+IMHT+1)%IMHT][(x+IMWD+1)%IMWD];
      if(A[y][x] == 1)
      {
        if (sum < 2)
           B[y][x] = 0;
        else if (sum == 2 || sum == 3)
           B[y][x] = 1;
        else
           B[y][x] = 0;
      }
      else
      {
        if (sum == 3)
          B[y][x] = 1;
        else
          B[y][x] = 0;
      }
    }
  }
}