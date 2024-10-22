#include <iostream>
using namespace std;

struct Node {
  int data;
  Node* prev;
  Node* next;
};

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

    Node* newNode = new Node{data, temp, temp->next};
    if (temp->next != nullptr) {
      temp->next->prev = newNode;
    }
    temp->next = newNode;
  }
  void printList() {
    Node* temp = head;
    cout << "Data yang dimasukkan ke dalam daftar adalah: " << endl;
    int i = 1;while (temp != nullptr) {
      cout << "node " << i++ << ": " << temp->data << endl;
      temp = temp->next;
    }
  }
};
int main() {
  DoubleLinkedList dll;
  int jumlahNode, data, posisi;
  cout << "Masukkan jumlah node (3 atau lebih): ";
  cin >> jumlahNode;

  for (int i = 1; i <= jumlahNode; i++) {
    cout << "Input data untuk node " << i << ": ";
    cin >> data;
    dll.append(data);
  }
  dll.printList();
  cout << endl;
  cout << "Masukkan posisi (2 hingga " << jumlahNode << ") untuk memasukkan node baru: ";
  cin >> posisi;
  cout << "Input data untuk posisi " << posisi << ": ";
  cin >> data;
  dll.insertAtPosition(data, posisi);
  cout << endl << "Setelah penyisipan, daftar baru adalah:" << endl;
  dll.printList();

  return 0;
}