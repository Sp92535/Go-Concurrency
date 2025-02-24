### **Experiment 7: Read-Write Locks (`sync.RWMutex`)**  

#### **Understanding Read-Write Locks in Go**  
A **race condition** can still happen if multiple goroutines read from a shared resource while another goroutine modifies it. A **read-write lock (`sync.RWMutex`)** allows:  
- **Multiple goroutines to read** simultaneously.  
- **Only one goroutine to write** at a time (blocking readers & writers).  

#### **Key Concepts**  

1. **Why Use `sync.RWMutex`?**  
   - If multiple goroutines **only read data**, locking with a normal `Mutex` would be inefficient.  
   - `sync.RWMutex` allows **concurrent reads** while ensuring **only one write at a time**.  

2. **How `RWMutex` Works:**  
   - `RLock()` → Multiple readers can proceed together.  
   - `RUnlock()` → Releases the read lock.  
   - `Lock()` → Blocks all readers and writers, ensuring exclusive write access.  
   - `Unlock()` → Releases the write lock.  

---

### **Problem Statement**  
Create a program where **multiple goroutines read** a shared variable while **one goroutine writes** to it periodically. Use `sync.RWMutex` to prevent data inconsistencies.  

---

### **Steps for Implementation**  

1. **Create a shared variable.**  
2. **Launch multiple reader goroutines that fetch data using `RLock()`.**  
3. **Launch a writer goroutine that updates data using `Lock()`.**  
4. **Observe incorrect output if no synchronization is used.**  
5. **Add `sync.RWMutex` to ensure readers don’t block each other but writers do.**  