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

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(nullptr);

  ll X, Q;
  cin >> X >> Q;

  multiset<ll> lower, upper;
  lower.insert(X);

  rep(i, Q) {
    ll A, B;
    cin >> A >> B;

    // 両方をlowerに追加
    lower.insert(A);
    lower.insert(B);

    // lowerの最大値2つをupperに移動
    upper.insert(*lower.rbegin());
    lower.erase(prev(lower.end()));
    upper.insert(*lower.rbegin());
    lower.erase(prev(lower.end()));

    // upperの最小値1つをlowerに移動（lowerが1つ多くなる）
    lower.insert(*upper.begin());
    upper.erase(upper.begin());

    // 中央値 = lowerの最大値
    cout << *lower.rbegin() << "\n";
  }

  return 0;
}
