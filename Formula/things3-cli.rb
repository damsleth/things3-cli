class Things3Cli < Formula
  desc "CLI for Things 3"
  homepage "https://github.com/ossianhempel/things3-cli"
  url "https://github.com/ossianhempel/things3-cli/archive/7a0a4447170e7a54ece9b33fce55d1dc6d696726.tar.gz"
  sha256 "9d503c8c8112d933946e3364ac6cb3ee90b9c86cc6bdd9644ee8fe64853efff2"
  version "20260102104112"

  depends_on "go" => :build

  def install
    ld_version = "7a0a444"
    ldflags = "-s -w -X github.com/ossianhempel/things3-cli/internal/cli.Version=#{ld_version}"
    system "go", "build", "-trimpath", "-ldflags", ldflags, "-o", bin/"things", "./cmd/things"
  end

  test do
    system "#{bin}/things", "--version"
  end
end
