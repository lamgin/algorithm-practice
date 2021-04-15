package com.algorithm.pattern.proxy;

import java.lang.annotation.Annotation;
import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;

public class JdkActorProxy<T>  {

    // 如果有需要代理的类
    private T obj;

    public Object getProxy(T obj){
        this.obj = obj;
        return  Proxy.newProxyInstance(this.getClass().getClassLoader(), new Class[]{Actor.class}, new ProxyHandler());
    }

    class ProxyHandler implements InvocationHandler{
        @Override
        public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
            System.out.println("before hello");
            System.out.println("==============");
            System.out.println("proxy:" + proxy.getClass());
            Annotation[] proxyAnnotations = proxy.getClass().getDeclaredAnnotations();
            System.out.println("annotation:");
            for (Annotation annotation : proxyAnnotations) {
                System.out.println(annotation.toString());
            }

            System.out.println("==============");
            Class clazz = method.getDeclaringClass();
            System.out.println("clazz:" + clazz.getClass());
            System.out.println("annotation:");
            Annotation[] clazzAnnotations = clazz.getDeclaredAnnotations();
            for (Annotation annotation : clazzAnnotations) {
                System.out.println(annotation.toString());
            }

            System.out.println("==============");
            System.out.println("method:" + method.getClass());
            System.out.println("generic:" + method.toGenericString());
            System.out.println("annotation:");
            Annotation[] annotations = method.getDeclaredAnnotations();
            for (Annotation annotation : annotations) {
                System.out.println(annotation.toString());
            }

            if(obj != null) {
                method.invoke(obj,args);
            }

            System.out.println("after hello");
            return null;
        }
    }

    /**
     * JDK 动态代理执行模式
     * @param args
     */
    public static void main(String[] args) {
        Actor actor1 = (Actor) (new JdkActorProxy<Actor>()).getProxy(new RealActor());
        actor1.doSomething();
        System.out.println("----------------------");
        Actor actor2 = (Actor) (new JdkActorProxy<Actor>()).getProxy(null);
        actor2.doSomething();
    }
}
