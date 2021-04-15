package com.algorithm.pattern.proxy;

public class RealActor implements Actor {
    @Override
    public void doSomething() {
        System.out.println("really doing something");
    }
}
