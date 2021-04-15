package com.algorithm.pattern.proxy;

@AnnotationActorClass(value = "actor")
public interface Actor {
    @AnnotationActorMethod(methodValue = "methodValue")
    @AnotherAnnotationActorMethod
    void doSomething();
}
