# Mid-term test (variant 3)

Complete the tasks below. Commit and push changes to the generated repository.
You don't need to create a pull request.

## 1. Save the world

You are a young computer scientist who is assigned to maintain an ancient
computation complex Wahnsinn-X. It is still being used by military forces to
predict and prevent a nuclear threat on the planet. You don't even want to think
about possible consequences of the "prevention".

It is your regular maintenance procedure, when you noticed a cat suspiciosly
looking at a cup of coffee forgotten by your colleague earlier this morning. The
cat pushes the cup with their paw. The cup falls down and the remaining coffee
splashes on the motherboard of Wanshinn-X. The next moment you hear the sirens,
see the red lights and the armored nuclear-proof shields seal the room closed.

You grab your phone and quickly create a filtering bridge between Wahsinn-X and
the upstream system. You always keep your favorite code editor at your finger
tips so it doesn't take any time to setup your favorite coding environment and
start observing.

The complex consists of various sensors that generate positive integer numbers.
But because of the short circuit, Whansinn-X has started to generate false
signals which seem to be always odd numbers followed by 4 or 2 in the output.
E.g. `34` or `72` actually should be just `4` or `2`.

Write a program in the folder `filter-digits` that accepts integer numbers,
filters out invalid digits and outputs the correct numbers. Since you are a great
engineer, writing tested code and ensuring all representative cases are covered
by the tests is a must. Tests guarantee you program correctness and the world safety.

You don't have much time left! The future of the world is in your hands!

### Example 1

```
Enter Wanshinn-X sensors sequence: 123456789
Correct is sequence is           : 2456789
```

Odd numbers `1` and `3` are followed by `4` or `2`, so they have to be removed.

### Example 2

```
Enter Wanshinn-X sensors sequence: 13579
Correct is sequence is           : 13579
```

## 2. Hack the cypher

You have implemented the filter and restored the communication flow of
Wahnsinn-X complex. However the nuclear thread prevention sequence has been
already launched and you can see red numbers on a black dashboard going down.
You understand it being a humanity final countdown.

You remember hearing about the Doof Maulwurf system that can prevent a doom day.
And you are very lucky to find a gateway to Doof Maulwurf through the Dark Net.
Hackers of the whole world are trying to break the code of Maulwurf. 

After heavy brainstorming with other hackers and multiple hacking attempts you
are sure that you have cracked the sequence that would enable Doof Maulwurf and
save the world:

- In row 0 (the topmost row), there is a unique nonzero entry `1`.
- Each entry of each subsequent row is constructed by adding the number above
  and to the left with the number above and to the right, treating blank entries
  as `0`.
  
For example, the initial number in the first (or any other) row is `1` (the sum
of `0` and `1`), whereas the numbers `1` and `3` in the third row are added to
produce the number `4` in the fourth row:

![triangle animation](https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif)

Or in an array form:
```
1 0 0 0 0
1 1 0 0 0
1 2 1 0 0
1 3 3 1 0
1 4 6 4 1
```

Your task is to create an application called `doof-maulwurf` which accepts a
positive number of rows of the key to generate, and outputs the key. Since it's
important the program works correct, you also have to create tests for it.

Hurry up! Doof Maulwurf must be activated to save the humanity!

### Example 1

```
The program generates a key to enable Doof Maulwurf.

Enter amount of rows to generate: 1

The key is:
1
```

### Example 2

```
The program generates a key to enable Doof Maulwurf.

Enter amount of rows to generate: 5

The key is:
1 0 0 0 0
1 1 0 0 0
1 2 1 0 0
1 3 3 1 0
1 4 6 4 1
```
