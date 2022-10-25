using System;

namespace csharp {
	class CSharp {
		static void Main(string[] args) {
			Random random = new Random();
			string password = "random" + random.Next();
		}
	}
}
