#include <iostream>
#include <vector>
#include "queue"

using namespace std;

void bfs(int n, int s, vector< vector<int> > g, vector<int>& d) {
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
                d[to] = d[v] + 1;
            }
        }
    }
}

int main() {
    int n, m, k;
    cin >> n >> m;
    vector<vector<int>> g(n);
    for (int i=0; i<m; ++i) {
        int u, v;
        cin >> u >> v;
        g[u].push_back(v);
        g[v].push_back(u);
    }
    cin >> k;
    vector<int> pivot;
    vector<bool> used (n);
    for (int i=0; i<k; ++i) {
        int p;
        cin >> p;
        pivot.push_back(p);
    }
    int count = 0;
    for (int i=0; i<n; ++i) {
        int temp = -1;
        bool is_pivot = true;
        vector<int> d(n);
        bfs(n, i, g, d);
        for (auto p : pivot) {
            if (temp == -1) {
                temp = d[p];
            }
            if (d[p] == 0 || i==p || temp != d[p]) {
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
