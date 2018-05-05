using System;
using Quobject.SocketIoClientDotNet.Client;

namespace NetConsoleApp
{
    class Program
    {
        static void Main(string[] args)
        {
            var socket = IO.Socket("http://localhost:5000");
            socket.On(Socket.EVENT_CONNECT, () =>
            {
                Console.WriteLine("Connected");
                socket.Emit("chat message", "Hi From .NET");
            });

            socket.On(Socket.EVENT_DISCONNECT, () =>
            {
                Console.WriteLine("Disconnected");
            });

            socket.On(Socket.EVENT_ERROR, data =>
            {
                Console.WriteLine($"Error: {data}");
            });

            socket.On("chat message", data =>
            {
                Console.WriteLine($"Receive: {data}");
            });

            Console.WriteLine("Enter to exit.");
            Console.ReadLine();

            socket.Disconnect();
        }
    }
}
