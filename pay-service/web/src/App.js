
import './App.css';
import OrderPage from './pages/OrderPage';
import ErrorPage from './pages/ErrorPage';
import { BrowserRouter , Route, Link, NavLink,Routes } from "react-router-dom";


function App() {
    return (

    <BrowserRouter>
        <Routes>
             <Route path="/order/:orderId" element={<OrderPage />} exact/>
             <Route path="/*" element={<ErrorPage />} exact/>
        </Routes>
    </BrowserRouter>

    );
}
export default App;

