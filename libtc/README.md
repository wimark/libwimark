# Docs for libtc

## Structure of tc objects

Objects consist on qdiscs (all with htb type), filters (all u32) and classes.
Qdiscs are:
 - root qdisc with id 0x8001
 - leaf qdiscs with ids from 0x1 to 0x1000 (created by demand)
Root qdisc has classes from 0x1 to 0x1000, each has attached leaf qdisc with the same id
Leaf qdiscs have up to 8 classes for each user, with ids from N*8+1 to N*8+8, where N is 12 bits of MAC
Each qdisc has 3 levels of filters:
 - first level has 1 filter redirecting to table 0x7ff bucket N, where N is 8 bits of MAC
 - second level has table 0x7ff with up to 256 buckets with 1 filter in each, redirecting to table N bucket M, where M is another 4 bits of MAC
 - third level has tables from 0x0 to 0xff with up to 256 buckets with filters, redirectind to class:
  - MN+1 for root qdisc (only 1 filter with id 0)
  - MN*8+K for leaf qdisc, where K is number of QoS for user class (1 to 8) (number of filters defined by traffic classes)
First filter level of root qdisc also has high-priority filter for unauthorized users, redirecting to default class 0xffff with rates for unauth users.

Lets user MAC be 01:23:45:67:89:AB. Then traffic flow for this user may be following:
 - root qdisc 8001
  - l1 filter 800:0:1 (id is default) -> table 7ff hash AB
  - l2 filter 7ff:AB:0 -> table AB hash 9
  - l3 filter AB:9:0 -> class 8001:9AB+1
 - leaf qdisc 9AB+1
  - filter 800:0:1 (id is default) -> table 7ff hash 67
  - filter 7ff:67:0 -> table 67 hash 8
  - filter 67:8:1 -> class 9AB+1:678*8+1 OR
  - filter 67:8:2 -> class 9AB+1:678*8+1 OR
  - filter 67:8:3 -> class 9AB+1:678*8+2 OR
  - filter 67:8:4 -> class 9AB+1:678*8+3 ...

### Initial

Created on interface init.

 - root qdisc: id 8001, parent root, default class ffff
 - filter table 7ff, parent 8001
 - 1st level filter: parent 8001, default id, priority 100, link to table 7ff by hash 0x0000 0000 00ff

 - unauth filter: parent 8001, default id, priority 1, match mark 777, to class ffff
 - unauth class: id ffff, rate from default config

### Per user

Created on user auth. Lets user MAC be 01:23:45:67:89:AB. Numbers in brackets are formed from MAC parts. 

 - filter table (AB), parent 8001
 - 2nd level filter: parent 8001, id 7ff:(AB):0, link to table (AB) by hash 0x0000 0000 0f00
 - root class: id 8001:(9AB)+1, rate unlimited
 - 3rd level filter: parent 8001, id (AB):(9):0, link to class 8001:(9AB)+1

 - leaf qdisc: id (9AB)+1, parent 8001:(9AB)+1, default class ffff
 - filter table 7ff, parent (9AB)+1
 - 4th level filter: parent (9AB)+1, default id, priority 100, link to table 7ff by hash 0x0000 00ff 0000

 - filter table (67), parent (9AB)+1
 - 5th level filter: parent (9AB)+1, id 7ff:(67):0, link to table (67) by hash 0x0000 0000 f000

 - qos1 class: id (9AB)+1:(678)*8+1, rate from qos1
 - 6th level qos1 filter1: parent (9AB)+1, id (67):(8):1, link to class (9AB)+1:(678)*8+1 match constraint1 from qos1
 - 6th level qos1 filter2: parent (9AB)+1, id (67):(8):2, link to class (9AB)+1:(678)*8+1 match constraint2 from qos1
 - ...
 - 6th level qos1 filter2: parent (9AB)+1, id (67):(8):N, link to class (9AB)+1:(678)*8+1 match constraintN from qos1

 - qos2 class: id (9AB)+1:(678)*8+2, rate from qos2
 - 6th level qos2 filter1: parent (9AB)+1, id (67):(8):N+1, link to class (9AB)+1:(678)*8+2 match constraint1 from qos2
 - 6th level qos2 filter2: parent (9AB)+1, id (67):(8):N+2, link to class (9AB)+1:(678)*8+2 match constraint2 from qos2
 - ...

### Constraints

 * classes and filters are matched consequently so order does matter
 * default class for user has no limits, so you better define last no-constraint QoS WITH limits for each user group
 * maximum number of QoS classes is 8
 * maximum number of filters is 4095 for all classes combined
 * users whose last 3 octets of MAC are the same, will share QoS buckets (for now)
