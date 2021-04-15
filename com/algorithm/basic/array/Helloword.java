package com.algorithm.basic.array;

import java.util.function.Supplier;

class BasicArray {
    interface Testor<T> {
         T test();
    }

    public static void main(String[] args) {
        Testor<BasicArray> obj = BasicArray::new;
        Object objj = obj;
        System.out.println(objj.getClass().toString());
        System.out.println(obj.getClass().getSuperclass().toString());
        if(obj instanceof Object){
            System.out.println("hei");
            Object anotherObj = (Object) obj;
            System.out.println(anotherObj);
        }
        Object newItem = obj.test();
        System.out.println(newItem.toString());
    }
}