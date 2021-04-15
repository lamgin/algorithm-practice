package com.algorithm.pattern.proxy;

public class ActorProxy implements Actor{

    Actor actor = new RealActor();

    @Override
    public void doSomething() {
        System.out.println("before act");
        actor.doSomething();
        System.out.println("after act");
    }

    /**
     * 静态代理执行模式
     * @param args
     */
    public static void main(String[] args) {
        (new ActorProxy()).doSomething();
    }
}
