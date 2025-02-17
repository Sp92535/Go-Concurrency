### **Experiment 1: Goroutines and Basic Concurrency**  

#### **Understanding Concurrency in Go**  
Go provides built-in support for concurrency using **goroutines**, which are lightweight threads managed by the Go runtime. Unlike traditional OS threads, goroutines have minimal overhead, allowing thousands to run simultaneously.  

#### **Key Concepts**  

1. **Goroutines**  
   - A goroutine is created using the `go` keyword before a function call.  
   - They execute independently and may not run in a sequential order.  
   - The main function does not wait for goroutines to complete by default.  

2. **Concurrency vs. Parallelism**  
   - **Concurrency** means multiple tasks start, run, and complete in overlapping time periods.  
   - **Parallelism** means tasks are executed literally at the same time using multiple CPU cores.  
   - Go supports concurrency by default, and parallelism is controlled using `runtime.GOMAXPROCS()`.  

3. **Goroutine Scheduling**  
   - Go uses **M:N scheduling**, meaning `M` goroutines are mapped to `N` OS threads dynamically.  
   - The Go runtime decides how to schedule goroutines, and their execution order is not guaranteed.  

4. **Main Function and Goroutines**  
   - If the main function exits, all goroutines stop, even if they haven't completed execution.  
   - A delay (like `time.Sleep()`) may be needed to let goroutines finish before `main` exits.  
   - A better approach is using **synchronization mechanisms** like `sync.WaitGroup`.  

---

### **Problem Statement**  
Create a program where multiple goroutines print numbers from different ranges concurrently. The goal is to observe how goroutines execute independently and how the main function's execution affects them.

---

### **Steps for Implementation**  

1. **Create a function that prints numbers in a loop.**  
2. **Call that function normally first (without concurrency) to observe sequential execution.**  
3. **Convert the function into a goroutine using the `go` keyword and observe execution order.**  
4. **Notice that the main function exits before goroutines complete.**  
5. **Introduce a delay (`time.Sleep()`) to allow goroutines to run.**  
6. **Recognize that `time.Sleep()` is not a proper way to synchronize goroutines.**  
7. **Learn about `sync.WaitGroup` and how it properly synchronizes goroutines.**  
8. **Analyze how goroutines behave when executing multiple concurrent tasks.**  