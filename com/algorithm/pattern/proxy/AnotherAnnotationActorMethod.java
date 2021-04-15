package com.algorithm.pattern.proxy;

import java.lang.annotation.*;

@Documented
@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.METHOD)
public @interface AnotherAnnotationActorMethod {
}
