#include <iostream>
#include <vector>
#include "queue"

using namespace std;

void bfs(int n, int s, vector< vector<int> > g, vector<vector<int>>& d, int j) {
    queue<int> q;
    q.push (s);
    vector<bool> used (n);
    used[s] = true;
    while (!q.empty()) {
        int v = q.front();
        q.pop();
        for (size_t i=0; i<g[v].size(); ++i) {
            int to = g[v][i];
            if (!used[to]) {
                used[to] = true;
                q.push (to);
                d[j][to] = d[j][v] + 1;
            }
        }
    }
}

int main() {
    int n, m, k;
    cin >> n >> m;
    vector<vector<int>> g(n);
    for (int i = 0; i < m; ++i) {
        int u, v;
        cin >> u >> v;
        g[u].push_back(v);
        g[v].push_back(u);
    }
    cin >> k;
    vector<int> pivot;
    for (int i = 0; i < k; ++i) {
        int p;
        cin >> p;
        pivot.push_back(p);
    }
    vector<vector<int>> d(k, vector<int>(n));
    int count = 0;
    for (int i=0; i<k; ++i) {
        bfs(n, pivot[i], g, d, i);
    }
    for (int i=0; i<n; ++i) {
        bool is_pivot = true;
        for (int j=1; j<k; ++j) {
            if (d[j][i] != d[j-1][i] || d[j][i] == 0) {
                is_pivot = false;
                break;
            }
        }
        if (is_pivot) {
            cout << i << endl;
            ++count;
        }
    }
    if (count == 0) cout << "-" << endl;
    return 0;
}
