package com.github.saser.adventofcode.year2016.day12;

import java.io.FileReader;
import java.io.IOException;

import org.junit.Test;
import org.junit.Assert;

public class Day12Test {
    @Test
    public void part1Actual() throws IOException {
        try (var input = new FileReader("inputs/2016/12")) {
            var output = "318020";
            var result = Day12.part1(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }

     @Test
     public void part2Actual() throws IOException {
         try (var input = new FileReader("inputs/2016/12")) {
             var output = "9227674";
             var result = Day12.part2(input);
             Assert.assertEquals("no error", "", result.error);
             Assert.assertEquals("correct output", output, result.answer);
         }
     }
}
