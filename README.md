# walls
Firewall with TUI using netfilter-queue Golang bindings.

### To run the project
```bash
# Clone this repository
git clone https://github.com/SaurusXI/protecc.git

# Setup iptable rules to route all incoming packets to NFQUEUE
sudo iptables -A INPUT -j NFQUEUE --queue-num 0 
# OR the following, if you prefer a non-blocking packet queue
sudo iptables -A INPUT -j NFQUEUE --queue-num 0 --queue-bypass

# Run the firewall to listen on packets on NFQUEUE
sudo go run src/main.go

# To remove the iptables rule (optional) run
sudo iptables -D INPUT -j NFQUEUE --queue-num 0
```
