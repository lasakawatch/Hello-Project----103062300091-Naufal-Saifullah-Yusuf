package main

import "fmt"

func main() {
	fmt.Println("================================")
	fmt.Println("         Soal 2. Faktor Bilangan")
	fmt.Println("               Disusun oleh:")
	fmt.Println("        Naufal Saifullah Yusuf")
	fmt.Println("================================")

	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scanln(&n)

	fmt.Print("Faktor-faktor dari ", n, ": ")
	cetakFaktorRekursif(n, 1)
}

func cetakFaktorRekursif(n, i int) {
	if i <= n {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
		cetakFaktorRekursif(n, i+1)
	}
}

class DoubleLinkedList {
 private: Node* head;

 public:
  DoubleLinkedList() : head(nullptr) {}

  void append(int data) {
    Node* newNode = new Node{data, nullptr, nullptr};
    if (head == nullptr) {head = newNode;return; }

    Node* temp = head;
    while (temp->next != nullptr) { temp = temp->next;}
    temp->next = newNode; newNode->prev = temp;
  }
  void insertAtPosition(int data, int position) {
    if (position <= 1) {
      cout << "Posisi tidak valid. Memasukkan di awal daftar.\n";
      Node* newNode = new Node{data, nullptr, head};
      if (head != nullptr) {head->prev = newNode;}
      head = newNode;return;
    }

    Node* temp = head;
    int count = 1;
    while (temp != nullptr && count < position - 1) {
      temp = temp->next;count++;
    }
    if (temp == nullptr) {
    cout << "Posisi tidak valid. Memasukkan di akhir daftar.\n";
      append(data);return;
    }
class DoubleLinkedList {
 private: Node* head;

 public:
  DoubleLinkedList() : head(nullptr) {}

  void append(int data) {
    Node* newNode = new Node{data, nullptr, nullptr};
    if (head == nullptr) {head = newNode;return; }

    Node* temp = head;
    while (temp->next != nullptr) { temp = temp->next;}
    temp->next = newNode; newNode->prev = temp;
  }
  void insertAtPosition(int data, int position) {
    if (position <= 1) {
      cout << "Posisi tidak valid. Memasukkan di awal daftar.\n";
      Node* newNode = new Node{data, nullptr, head};
      if (head != nullptr) {head->prev = newNode;}
      head = newNode;return;
    }

    Node* temp = head;
    int count = 1;
    while (temp != nullptr && count < position - 1) {
      temp = temp->next;count++;
    }
    if (temp == nullptr) {
    cout << "Posisi tidak valid. Memasukkan di akhir daftar.\n";
      append(data);return;
    }
