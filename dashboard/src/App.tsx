import './App.css';
import { ViewsByTime } from './components/Graphs';
import { TopPagesTable, TopReferersTable } from './components/Tables';

function App() {
  return (
    <div className="app">
      <h1>Dashboard</h1>

      <div>
        <ViewsByTime />
      </div>

      <div>
        <div className="two-column">
          <TopPagesTable />
          <TopReferersTable />
        </div>
      </div>
    </div>
  );
}

export default App;
