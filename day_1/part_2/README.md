## Day 1 - Part 2
---
You notice that the device repeats the same frequency change list over and over. To calibrate the device, you need to find the first frequency it reaches twice.

For example, using the same list of changes above, the device would loop as follows:
```
Current frequency  0, change of +1; resulting frequency  1.
Current frequency  1, change of -2; resulting frequency -1.
Current frequency -1, change of +3; resulting frequency  2.
Current frequency  2, change of +1; resulting frequency  3.
```
(At this point, the device continues from the start of the list.)
```
Current frequency  3, change of +1; resulting frequency  4.
Current frequency  4, change of -2; resulting frequency  2, which has already been seen.
```
In this example, the first frequency reached twice is 2. Note that your device might need to repeat its list of frequency changes many times before a duplicate frequency is found, and that duplicates might be found while in the middle of processing the list.


Here are other examples:
```
+1, -1 first reaches 0 twice.
+3, +3, +4, -2, -4 first reaches 10 twice.
-6, +3, +8, +5, -6 first reaches 5 twice.
+7, +7, -2, -7, -4 first reaches 14 twice.
```
What is the first frequency your device reaches twice?
---
### Some Performance Metrics per Round
```
1: 125.233µs
2: 307.155µs
3: 369.927µs
4: 344.592µs
5: 337.913µs
6: 275.117µs
7: 351.54µs
8: 411.876µs
9: 710.018µs
10: 741.725µs
11: 790.052µs
12: 929.711µs
13: 959.762µs
14: 1.096262ms
15: 1.171164ms
16: 1.169249ms
17: 1.21995ms
18: 1.30363ms
19: 1.505555ms
20: 1.466069ms
21: 1.496965ms
22: 1.53358ms
23: 1.81271ms
24: 2.149716ms
25: 1.888458ms
26: 1.895177ms
27: 2.094521ms
28: 1.94251ms
29: 1.968575ms
30: 2.250511ms
31: 2.395268ms
32: 2.636985ms
33: 2.33217ms
34: 2.352787ms
35: 2.363516ms
36: 2.430585ms
37: 2.536714ms
38: 3.714679ms
39: 2.718306ms
40: 2.799553ms
41: 2.774763ms
42: 2.852219ms
43: 3.234863ms
44: 3.552434ms
45: 3.340428ms
46: 3.151816ms
47: 3.231582ms
48: 4.334333ms
49: 3.731611ms
50: 3.724624ms
51: 3.676026ms
52: 3.76222ms
53: 4.600311ms
54: 4.033613ms
55: 3.791601ms
56: 3.98131ms
57: 4.841553ms
58: 4.182366ms
59: 4.169551ms
60: 5.690464ms
61: 5.086825ms
62: 4.886788ms
63: 4.957608ms
64: 5.672268ms
65: 5.202135ms
66: 5.29539ms
67: 6.129791ms
68: 5.529853ms
69: 5.821095ms
70: 6.704824ms
71: 5.753055ms
72: 6.912737ms
73: 6.28793ms
74: 6.123946ms
75: 8.913549ms
76: 6.433144ms
77: 7.633369ms
78: 6.879643ms
79: 8.007787ms
80: 7.335142ms
81: 7.648625ms
82: 7.88587ms
83: 7.700588ms
84: 8.503053ms
85: 7.874817ms
86: 9.012721ms
87: 8.155664ms
88: 8.930665ms
89: 8.660821ms
90: 8.713721ms
91: 9.291008ms
92: 8.884294ms
93: 9.540799ms
94: 9.100759ms
95: 10.184952ms
96: 8.813813ms
97: 10.305193ms
98: 9.928659ms
99: 9.282891ms
100: 10.383376ms
101: 9.338145ms
102: 10.70206ms
103: 10.33435ms
104: 9.92613ms
105: 11.017543ms
106: 10.145579ms
107: 11.026825ms
108: 11.353707ms
109: 10.64697ms
110: 11.716911ms
111: 12.576524ms
112: 11.78352ms
113: 10.946904ms
114: 12.428068ms
115: 12.11343ms
116: 11.92873ms
117: 11.907683ms
118: 12.878451ms
119: 12.180247ms
120: 11.498571ms
121: 12.418244ms
122: 12.660231ms
123: 12.656235ms
124: 11.648493ms
125: 12.78418ms
126: 13.108834ms
127: 12.728022ms
128: 11.954906ms
129: 12.879807ms
130: 13.380252ms
131: 13.550415ms
132: 13.685677ms
133: 12.530445ms
134: 14.138294ms
135: 14.666025ms
136: 14.209613ms
137: 14.304277ms
Found repeating frequency: 66105
Found in 137 interations!
```