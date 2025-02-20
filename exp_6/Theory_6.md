### **Experiment 6: Mutex and Race Conditions**  

#### **Understanding Race Conditions in Go**  
A **race condition** occurs when multiple goroutines access and modify the same shared variable **without synchronization**, leading to unpredictable results.  

#### **Key Concepts**  

1. **Why Race Conditions Happen**  
   - When multiple goroutines read and write a shared variable simultaneously, **one update may override another**, causing data corruption.  
   - Since goroutines execute **independently**, the final value depends on unpredictable execution order.  

2. **Detecting Race Conditions**  
   - Go provides a built-in **race detector**:  
     ```sh
     go run -race main.go
     ```
   - Running this will **flag race conditions** in the output.  

3. **Fixing Race Conditions Using `sync.Mutex`**  
   - A **mutex (mutual exclusion lock)** ensures only **one goroutine modifies a shared variable at a time**.  
   - `sync.Mutex` provides:  
     - `Lock()` – Blocks other goroutines from accessing the resource.  
     - `Unlock()` – Releases the lock, allowing other goroutines to proceed.  

---

### **Problem Statement**  
Create a **shared counter** that multiple goroutines increment concurrently, leading to a **race condition**. Then, **fix the race condition** using `sync.Mutex`.  

---

### **Steps for Implementation**  

1. **Create a global integer counter.**  
2. **Launch multiple goroutines to increment the counter simultaneously.**  
3. **Run the program with the `-race` flag and observe inconsistent results.**  
4. **Introduce `sync.Mutex` to lock and unlock the counter during updates.**  
5. **Run the program again to confirm the race condition is fixed.**  