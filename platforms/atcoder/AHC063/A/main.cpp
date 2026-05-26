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

struct Snake {
  vector<pii> pos;
  vector<int> color;
};


pii nextPos(pii pos, char dir) {
  if (dir == 'U') return {pos.first -1, pos.second};
  if (dir == 'D') return {pos.first +1, pos.second};
  if (dir == 'L') return {pos.first, pos.second -1};
  return {pos.first, pos.second +1};
};

bool canMove(Snake& snake, char dir, int N) {
  pii next = nextPos(snake.pos[0], dir);
  if (next.first < 0 || N <= next.first) return false;
  if (next.second < 0 || N <= next.second) return false;
  if (next == snake.pos[1]) return false;
  return true;
};

void move(Snake& snake, char dir, vector<vector<int>>& board, string& answer) {
  pii next = nextPos(snake.pos[0], dir);
  pii tail = snake.pos.back();
  (void)snake.color.back(); // tailColor unused for now

  // 座標更新（基本移動：先頭に追加、末尾を削除）
  snake.pos.insert(snake.pos.begin(), next);
  snake.pos.pop_back();

  // 食事判定
  int food = board[next.first][next.second];
  if (food > 0) {
    snake.pos.push_back(tail);       // 末尾を戻す（伸びる）
    snake.color.push_back(food);     // 色を末尾に追加
    board[next.first][next.second] = 0;
  }

  // 噛みちぎり判定（胴体 pos[1] ~ pos[k-2] と一致するか）
  int k = snake.pos.size();
  for (int h = 1; h <= k - 2; h++) {
    if (next == snake.pos[h]) {
      // h+1 以降を切り離して餌にする
      for (int p = h + 1; p < k; p++) {
        pii dropPos = snake.pos[p];
        board[dropPos.first][dropPos.second] = snake.color[p];
      }
      // 蛇を h+1 の長さに縮める
      snake.pos.resize(h + 1);
      snake.color.resize(h + 1);
      break;
    }
  }

  answer += dir;
}

vector<pii> findFood(int targetColor, vector<vector<int>>& board, int N) {
  vector<pii> result;
  rep(i, N) rep(j, N) {
    if (board[i][j] == targetColor) {
      result.push_back({i, j});
    }
  }
  return result;
}

vector<char> pathTo(Snake& snake, pii target, int N) {
  pii start = snake.pos[0];
  if (start == target) return {};

  vector<vector<int>> dist(N, vector<int>(N, -1));
  vector<vector<char>> from(N, vector<char>(N, ' '));
  queue<pii> q;
  q.push(start);
  dist[start.first][start.second] = 0;

  string dirs = "UDLR";
  while (!q.empty()) {
    pii cur = q.front(); q.pop();
    for (char d : dirs) {
      pii nxt = nextPos(cur, d);
      if (nxt.first < 0 || N <= nxt.first) continue;
      if (nxt.second < 0 || N <= nxt.second) continue;
      if (dist[nxt.first][nxt.second] >= 0) continue;
      dist[nxt.first][nxt.second] = dist[cur.first][cur.second] + 1;
      from[nxt.first][nxt.second] = d;
      q.push(nxt);
      if (nxt == target) break;
    }
  }

  vector<char> path;
  pii cur = target;
  while (cur != start) {
    char d = from[cur.first][cur.second];
    path.push_back(d);
    if (d == 'U') cur.first++;
    else if (d == 'D') cur.first--;
    else if (d == 'L') cur.second++;
    else cur.second--;
  }
  reverse(all(path));
  return path;
}

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(nullptr);

  int N, M, C;
  cin >> N >> M >> C;

  vector<int> target(M);
  rep(i, M) cin >> target[i];

  vector<vector<int>> board(N, vector<int>(N));
  rep(i, N) rep(j, N) cin >> board[i][j];

  string answer;

  // 蛇の初期化
  Snake snake;
  snake.pos = {{4,0}, {3,0}, {2,0}, {1,0}, {0,0}};
  snake.color = {1, 1, 1, 1, 1};

  // 貪欲法: target[5] から target[M-1] まで順に正しい色の餌を食べる
  for (int i = 5; i < M; i++) {
    int needColor = target[i];

    // その色の餌を探す
    vector<pii> foods = findFood(needColor, board, N);
    if (foods.empty()) continue; // 見つからない場合はスキップ

    // 最も近い餌を選ぶ
    pii head = snake.pos[0];
    pii nearest = foods[0];
    int minDist = abs(head.first - foods[0].first) + abs(head.second - foods[0].second);
    for (auto& f : foods) {
      int d = abs(head.first - f.first) + abs(head.second - f.second);
      if (d < minDist) {
        minDist = d;
        nearest = f;
      }
    }

    // 経路を計算して移動
    vector<char> path = pathTo(snake, nearest, N);
    for (char d : path) {
      if (canMove(snake, d, N)) {
        move(snake, d, board, answer);
      }
    }
  }

  // 出力
  for (char c : answer) {
    cout << c << '\n';
  }

  return 0;
}
