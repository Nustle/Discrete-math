#include <iostream>
#include <vector>
#include <map>
#include <string>
#include <tuple>

using namespace std;

struct couple {
    int f;
    string s;

    bool operator<(const couple &other) const {
        return tie(f, s) < tie(other.f, other.s);
    }
};

int main() {
    int m, q, n;
    cin >> m;
    vector<string> a(m);
    for (int i = 0; i < m; i++) {
        cin >> a[i];
    }

    cin >> q;
    vector<string> b(q);
    for (int i = 0; i < q; i++) {
        cin >> b[i];
    }

    cin >> n;
    vector<vector<int>> arr(n, vector<int>(m));
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++) {
            cin >> arr[i][j];
        }
    }

    vector<vector<string>> s(n, vector<string>(m));
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++) {
            cin >> s[i][j];
        }
    }

    vector<couple> pts;
    map<couple, int> nums;
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++) {
            int nm = stoi(s[i][j]);
            couple tuple = {arr[i][j], b[nm]};
            if (!nums.count(tuple)) {
                pts.push_back(tuple);
                nums[tuple] = pts.size() - 1;
            }
        }
    }

    cout << "digraph {\n    rankdir = LR\n";
    for (auto &pair : nums) {
        cout << "    " << pair.second << " [label = \"(" << pair.first.f << "," << pair.first.s << ")\"]\n";
    }

    for (auto &pair : nums) {
        for (int j = 0; j < m; j++) {
            int nm = stoi(s[pair.first.f][j]);
            cout << "    " << pair.second << " -> " << nums[{arr[pair.first.f][j], b[nm]}] << " [label = \"" << a[j] << "\"]\n";
        }
    }

    cout << "}\n";

    return 0;
}
