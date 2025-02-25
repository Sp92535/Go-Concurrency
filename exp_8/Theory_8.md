### **Experiment 8: Worker Pool Pattern**  

#### **Understanding the Worker Pool Pattern**  
A **worker pool** is a concurrency pattern where a fixed number of worker goroutines process tasks from a shared queue. This improves efficiency by **limiting the number of active goroutines** instead of creating one per task.  

#### **Key Concepts**  

1. **Why Use a Worker Pool?**  
   - Creating **too many goroutines** can lead to high memory usage.  
   - A fixed number of workers **efficiently process tasks in parallel**.  

2. **How It Works**  
   - A **task queue (channel)** holds jobs.  
   - **Workers (goroutines)** pick up tasks from the queue.  
   - A **WaitGroup** ensures all tasks are completed before exiting.  

---

### **Problem Statement**  
Create a **worker pool** where a fixed number of goroutines process multiple tasks from a queue. The program should:  
- Accept **N** tasks.  
- Process them using **M worker goroutines**.  
- Print results as tasks are completed.  

---

### **Steps for Implementation**  

1. **Create a buffered channel to hold tasks.**  
2. **Launch `M` worker goroutines that fetch tasks from the channel.**  
3. **Send `N` tasks into the channel.**  
4. **Use `sync.WaitGroup` to wait until all workers finish.**  
5. **Observe how workers process tasks concurrently.**  