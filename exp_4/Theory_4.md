### **Experiment 4: Buffered vs Unbuffered Channels**  

#### **Understanding Buffered and Unbuffered Channels**  
Go channels can be **buffered** or **unbuffered**, and understanding their differences is key to writing efficient concurrent programs.  

1. **Unbuffered Channels (Default)**  
   - Block the sender **until the receiver is ready**.  
   - Ensure strict synchronization between goroutines.  
   - Example:  
     ```go
     ch := make(chan int) // Unbuffered channel
     ```  

2. **Buffered Channels**  
   - Allow sending multiple values **without immediate blocking**, up to a fixed capacity.  
   - The sender blocks **only when the buffer is full**.  
   - Example:  
     ```go
     ch := make(chan int, 3) // Buffered channel with capacity 3
     ```  

3. **Reading from a Buffered Channel**  
   - Messages stay in the buffer until received.  
   - If the buffer is empty, the receiver blocks.  

4. **Closing the Channel**  
   - The sender should **close the channel** when done sending.  
   - Receivers can detect closure using:  
     ```go
     val, ok := <-ch // 'ok' is false if the channel is closed
     ```  

---

### **Problem Statement**  
Demonstrate the difference between **buffered** and **unbuffered** channels by implementing a simple **message queue**. Observe how goroutines behave differently in each case.  

---

### **Steps for Implementation**  

1. **Create an unbuffered channel and launch goroutines to send data.**  
2. **Observe how senders block until the receiver reads the data.**  
3. **Modify the channel to be buffered and set its capacity to a fixed number.**  
4. **Observe how multiple sends happen before blocking occurs.**  
5. **Print execution order to understand the impact of buffering.**  