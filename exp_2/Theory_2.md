### **Experiment 2: Syncing Goroutines with WaitGroups**  

#### **Understanding Synchronization in Go**  
In the previous experiment, we saw that the main function exits before goroutines complete execution unless we introduce artificial delays using `time.Sleep()`. However, using `time.Sleep()` is unreliable because execution times can vary. A proper way to ensure goroutines complete before the program exits is by using **sync.WaitGroup**.  

#### **Key Concepts**  

1. **Why Synchronization is Needed**  
   - The Go runtime schedules goroutines independently, so execution order is not guaranteed.  
   - Without synchronization, the main function might terminate while goroutines are still running.  

2. **sync.WaitGroup**  
   - A **WaitGroup** allows us to wait for multiple goroutines to finish before proceeding.  
   - It has three key methods:  
     - `Add(n int)`: Increments the counter by `n` (number of goroutines).  
     - `Done()`: Decrements the counter (called inside each goroutine).  
     - `Wait()`: Blocks execution until the counter reaches zero.  

3. **Using sync.WaitGroup Correctly**  
   - `WaitGroup` should be passed by reference to goroutines to avoid unintended behavior.  
   - Each goroutine should call `wg.Done()` to signal completion.  
   - Calling `wg.Wait()` in the main function ensures all goroutines finish before continuing.  

---

### **Problem Statement**  
Modify the previous experiment to ensure all goroutines finish before the main function exits. Use **sync.WaitGroup** to synchronize execution instead of using `time.Sleep()`.  

---

### **Steps for Implementation**  

1. **Import the `sync` package and declare a `sync.WaitGroup`.**  
2. **Before starting goroutines, call `wg.Add(n)`, where `n` is the number of goroutines.**  
3. **Inside each goroutine, call `wg.Done()` when the function completes execution.**  
4. **In the main function, call `wg.Wait()` to ensure all goroutines finish before proceeding.**  
5. **Observe that goroutines now complete execution properly without artificial delays.**  