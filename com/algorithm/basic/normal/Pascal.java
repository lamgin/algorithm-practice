package com.algorithm.basic.normal;

/**
 * 杨辉三角
 */
public class Pascal {

    public void print(int level){

        long[] steps = new long[level];
        for (int i = 0; i<level; i++) {
            long[] his = steps.clone();
            for (int j = 0; j<=i; j++) {
                // 第一层 第二层 输出1, 边界输出1
                if(i == 1 || j == 0) {
                    steps[j] = 1;
                } else {
                    steps[j] = his[j] + his[j-1];
                }
                System.out.printf("%d ",steps[j]);
            }
            System.out.println();
        }
    }

    public static void main(String[] args) {
        new Pascal().print(10);
    }

}
