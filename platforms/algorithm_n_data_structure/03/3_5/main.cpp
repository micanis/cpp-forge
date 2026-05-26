#include <algorithm>
#include <bitset>
#include <climits>
#include <cmath>
#include <deque>
#include <iostream>
#include <map>
#include <numeric>
#include <queue>
#include <set>
#include <stack>
#include <string>
#include <vector>

using namespace std;

#define rep(i, n) for (int i = 0; i < (int)(n); i++)
#define REP(i, a, b) for (int i = (a); i < (int)(b); i++)
#define rrep(i, n) for (int i = (int)(n) - 1; i >= 0; i--)
#define all(x) (x).begin(), (x).end()

using ll = long long;
using pii = pair<int, int>;
using pll = pair<ll, ll>;

constexpr int INF = 1e9;
constexpr ll LINF = 1e18;
constexpr int MOD = 1e9 + 7;

template <class T>
bool chmin(T& a, const T& b) {
  return a > b ? a = b, true : false;
}
template <class T>
bool chmax(T& a, const T& b) {
  return a < b ? a = b, true : false;
}

#ifdef DEBUG
template <class A, class B>
ostream& operator<<(ostream& os, const pair<A, B>& p) {
  return os << "(" << p.first << ", " << p.second << ")";
}
template <class T, class = decltype(begin(declval<T>())),
          class = enable_if_t<!is_same<T, string>::value>>
ostream& operator<<(ostream& os, const T& c) {
  os << "[";
  bool first = true;
  for (const auto& v : c) {
    if (!first) os << ", ";
    os << v;
    first = false;
  }
  return os << "]";
}
#define dbg(x) cerr << #x << " = " << (x) << endl
#else
#define dbg(x)
#endif

vector<string> bubble_sort(int card_count, vector<string> cards) {
  int sorted_boundary = 0;
  bool has_remaining = true;

  while (has_remaining) {
    has_remaining = false;
    for (int j = card_count - 1; j > sorted_boundary; j--) {
      string card = cards[j];
      int card_num = card[1] - '0';
      string prev_card = cards[j - 1];
      int prev_card_num = prev_card[1] - '0';

      if (card_num < prev_card_num) {
        swap(cards[j], cards[j - 1]);
      }
      has_remaining = true;
    }
    sorted_boundary++;
  }

  return cards;
}

vector<string> selection_sort(int card_count, vector<string> cards) {
  int sorted_boundary = 0;
  bool has_remaining = true;

  while (has_remaining) {
    has_remaining = false;
    int min_value_index = sorted_boundary;
    for (int j = card_count - 1; j > sorted_boundary; j--) {
      string card = cards[j];
      int card_num = card[1] - '0';
      string min_card = cards[min_value_index];
      int min_card_num = min_card[1] - '0';
      if (card_num < min_card_num) {
        min_value_index = j;
      }
      has_remaining = true;
    }
    swap(cards[sorted_boundary], cards[min_value_index]);
    sorted_boundary++;
  }

  return cards;
}

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(nullptr);

  int N;
  cin >> N;

  vector<string> cards(N);
  rep(i, N) cin >> cards[i];

  vector<string> bubble_sorted_cards = bubble_sort(N, cards);
  vector<string> selection_sorted_cards = selection_sort(N, cards);

  bool is_selection_stable = false;
  if (bubble_sorted_cards == selection_sorted_cards) is_selection_stable = true;

  rep(i, N) {
    if (i == N - 1)
      cout << bubble_sorted_cards[i] << "\n";
    else
      cout << bubble_sorted_cards[i] << " ";
  }
  cout << "Stable" << "\n";
  rep(i, N) {
    if (i == N - 1)
      cout << selection_sorted_cards[i] << "\n";
    else
      cout << selection_sorted_cards[i] << " ";
  }
  cout << (is_selection_stable ? "Stable" : "Not stable") << "\n";

  return 0;
}
