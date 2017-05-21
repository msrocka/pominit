package main

const jtestText = `
package tests;

import org.junit.Assert;
import org.junit.Test;

public class ATest {

    @Test
    public void test() {
        System.out.println("running a test ...");
        Assert.assertTrue(4 == 2 << 1);
    }
}
`

const ktestText = `
package tests

import org.junit.Assert
import org.junit.Test

class ATest {

    @Test
    fun test() {
        println("running a test ...")
        Assert.assertTrue(4 == 2 shl 1)
    }
}
`

const kappText = `
package app

fun main(args: Array<String>) {
    println("Works!")
}
`
