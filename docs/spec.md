## 1. Mapping of Xunit Results
 
Xunit **Pass**, maps to Passed in Polarion, example:

```
<testcase classname="shotwell" name="startViaCommand">
    <system-out>I am stdout!!</system-out>
    <system-err>I am stderr!!!</system-err>
  <properties>
    <property name="polarion-testcase-id" value="EL-1"/>
  </properties>
</testcase>
```

Xunit **Failure**, maps to Failed in Polarion and the first 222 characters will be added to the result, example:

```
<testcase classname="shotwell" name="faq">
    <failure message="Test:6 Requires:['']
No packages defined for this test
Test:6 Testname:faq
 Change Dir: /mnt/tests/shotwell/. 
 Change Mod ./runtest.sh faq 
 Run Test: ./runtest.sh faq Watchdog: 600(s) 
 Warning: Failing Test since return code was (1), running [./runtest.sh faq]
 ...
 Result: FAIL
" type="failure"/>
    <system-out>I am stdout!!</system-out>
    <system-err>I am stderr!!!</system-err>
  <properties>
    <property name="polarion-testcase-id" value="EL-24"/>
  </properties>
</testcase>
```

Xunit **Error**, maps to Blocked in Polarion, and the first 222 characters will be added to the result, example:

```
<testcase classname="shotwell" name="BLOCKED TEST">
    <error message="error found in main, test could not be executed" type="error"/>
    <system-out>I am stdout!!</system-out>
    <system-err>I am stderr!!!</system-err>
  <properties>
    <property name="polarion-testcase-id" value="EL-42"/>
  </properties>
</testcase>
```

Xunit **Skipped**, maps to Waiting in Polarion by default, using the polarion-include-skipped option, the results can be ignored: 

```
<testcase classname="shotwell" name="SKIPPED TEST">
    <skipped message="Invalid OS Skipped - rhel7" type="skipped"/>
    <system-out>I am stdout!!</system-out>
    <system-err>I am stderr!!!</system-err>
  <properties>
    <property name="polarion-testcase-id" value="EL-41"/>
  </properties>
</testcase>
```

- username: rhevm3_machine
- password: polarion